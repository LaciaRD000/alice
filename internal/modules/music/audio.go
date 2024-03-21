package music

import youtube "github.com/knadh/go-get-youtube/youtube"

func DownloadMusic(id string) error {
	video, err := youtube.Get(id)
	if err != nil {
		return err
	}

	// download the video and write to file
	option := &youtube.Option{
		Rename: true, // rename file using video title
		Resume: true, // resume cancelled download
		Mp3:    true, // extract audio to MP3
	}
	err = video.Download(0, "temp/temp.mp3", option)
	if err != nil {
		return err
	}

	return nil
}
