package scan

import (
	"context"
	"fmt"
	"time"

	"github.com/dcrosby42/picfinder/fileinfo"
	picfinder_grpc "github.com/dcrosby42/picfinder/grpc"
)

func ScanAndSend(client picfinder_grpc.PicfinderClient, thishost string, dirname string, scanAll bool) error {
	infoC := make(chan fileinfo.FileInfo, 100)

	// Start the file scan generator:
	if scanAll {
		fmt.Printf("Scanning ALL files in %q...\n", dirname)
		go GenerateAllFiles(thishost, dirname, infoC)
	} else {
		fmt.Printf("Scanning media files in %q...\n", dirname)
		go GenerateMediaFiles(thishost, dirname, infoC)
	}

	// Collect all file infos into a list
	fileInfos := make([]fileinfo.FileInfo, 0, 10000)
	started := time.Now()
	for info := range infoC {
		fileInfos = append(fileInfos, info)
	}
	elapsed := time.Now().Sub(started)
	total := len(fileInfos)

	fmt.Printf("Elapsed: %s\n", elapsed)
	fmt.Printf("Total files to process: %d\n", total)

	// Now, reiterate the files and generate content hashes and send them to the API:
	fmt.Printf("Processing files and sending to server...\n")
	started = time.Now()
	statusEvery := 100
	errors := make([]error, 0)
	for i, fileInfo := range fileInfos {
		err := UpdateContentHash(&fileInfo)
		if err != nil {
			fmt.Printf("!!! ERR calculating hash for fileInfo=%s err=%s\n", fileInfo, err)
			errors = append(errors, err)
			continue
		}

		err = SendFileToApi(client, &fileInfo)
		if err != nil {
			fmt.Printf("!!! ERR sending fileInfo=%s err=%s\n", fileInfo, err)
			errors = append(errors, err)
			continue
		}
		if i%statusEvery == 0 {
			fmt.Printf("Processed %d of %d files (%.1f%% done), elapsed: %s\n", i, total, 100*float64(i)/float64(total), time.Now().Sub(started))
		}
	}
	fmt.Printf("Processed %d files (100%% done), elapsed: %s\n", total, time.Now().Sub(started))
	if len(errors) > 0 {
		fmt.Printf("THERE WERE %d ERRORS\n", len(errors))
	} else {
		fmt.Printf("No errors\n")
	}

	return nil
}

func UpdateContentHash(info *fileinfo.FileInfo) error {
	chash, lower32, err := HashFileContentSha256(info.PathString())
	if err != nil {
		return err
	}
	info.ContentHash = chash
	info.ContentHashLower32 = lower32
	return nil
}

func SendFileToApi(client picfinder_grpc.PicfinderClient, info *fileinfo.FileInfo) error {
	request := &picfinder_grpc.AddFileRequest{}
	request.FileInfo = fileinfo.ToGrpcFileInfo(*info)
	resp, err := client.AddFile(context.Background(), request)
	if err != nil {
		return err
	}
	if resp.Header == nil {
		return fmt.Errorf("api.AddFile() no response header!")
	}
	if resp.Header.Status != 0 {
		return fmt.Errorf("api.AddFile() responded non-0 status. Status=%d Message=%q FileId=%d UpdateAction=%s", resp.Header.Status, resp.Header.Message, resp.FileId, resp.UpdateAction)
	}
	// if resp.UpdateAction != string(fileinfo.UpdateAction_Insert) {
	// 	fmt.Printf("FileId=%d UpdateAction=%s\n", resp.FileId, resp.UpdateAction)
	// }

	return nil
}
