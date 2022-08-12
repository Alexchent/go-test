package main

import (
	scan "github.com/alexchen/test/ScanService"
	"time"
)

func main()  {
	go func() {
		for {
			scan.Start("/Volumes/Newsmy/AV")
			scan.Start("/Volumes/Newsmy/FC2-PPV")
			time.Sleep(10 * time.Minute)
	    }
		
	}()

	go func() {
		for {
			//scan.Start("/Volumes/Seagate-2T/AV")
		    scan.Start("/Volumes/Seagate-2T")
			time.Sleep(10 * time.Minute)
		}
	}()

    go func() {
		for {
			scan.Start("/Volumes/4T/AV")
			scan.Start("/Volumes/4T/FC2-PPV")
			time.Sleep(10 * time.Minute)
		}
	}()

	go func() {
		for {
			scan.Start("/Volumes/1T")
			//scan.Start("/Volumes/1T/AV")
			//scan.Start("/Volumes/1T/FC2-PPV")
			time.Sleep(10 * time.Minute)
		}
	}()

    for {
			scan.Start("/Users/chentao/Downloads/xunlei")
			time.Sleep(1 * time.Minute)
	}
    
}