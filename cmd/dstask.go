package main

import (
	"github.com/naggie/dstask"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		os.Args = append(os.Args, "next")
	}

	switch(os.Args[1]) {
		case "next":
			ts := dstask.LoadTaskSetFromDisk(dstask.NORMAL_STATUSES)
			ts.SortTaskList()
			ts.Display()
		case "add":
			ts := dstask.LoadTaskSetFromDisk(dstask.NORMAL_STATUSES)
			tl := dstask.ParseTaskLine(os.Args[2:])
			ts.AddTask(dstask.Task{
				WritePending: true,
				Status: dstask.STATUS_PENDING,
				Summary: tl.Text,
				Tags: tl.Tags,
				Project: tl.Project,
				Priority: tl.Priority,
			})
			ts.SaveToDisk()
		case "start":
		case "stop":
		case "done":
		case "context":
		case "modify":
		case "edit":
		case "describe":
		case "projects":
		case "day":
		case "week":
		case "import":
			ts := dstask.LoadTaskSetFromDisk(dstask.ALL_STATUSES)
			ts.ImportFromTaskwarrior()
			ts.SaveToDisk()
		case "help":
			dstask.Help()
			os.Exit(1)
		default:
			dstask.Help()
			os.Exit(1)
	}
}