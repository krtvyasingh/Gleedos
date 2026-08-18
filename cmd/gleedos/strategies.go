package main

type downloadStrategy struct {
	name string
	args []string
}

func buildDownloadStrategies(
	base []string,
	quality string,
	audioMode bool,
	url string,
) []downloadStrategy {
	if audioMode {
		return []downloadStrategy{
			{
				name: "mweb audio",
				args: append(
					append([]string{}, base...),
					"--extractor-args",
					"youtube:player_client=mweb",
					"-f",
					"ba/b",
					"--extract-audio",
					"--audio-format",
					"mp3",
					"--audio-quality",
					"0",
					url,
				),
			},
			{
				name: "web_safari audio",
				args: append(
					append([]string{}, base...),
					"--extractor-args",
					"youtube:player_client=web_safari",
					"-f",
					"ba/b",
					"--extract-audio",
					"--audio-format",
					"mp3",
					"--audio-quality",
					"0",
					url,
				),
			},
			{
				name: "android audio",
				args: append(
					append([]string{}, base...),
					"--extractor-args",
					"youtube:player_client=android",
					"-f",
					"ba/b",
					"--extract-audio",
					"--audio-format",
					"mp3",
					"--audio-quality",
					"0",
					url,
				),
			},
		}
	}

	return []downloadStrategy{
		{
			name: "mweb best",
			args: append(
				append([]string{}, base...),
				"--extractor-args",
				"youtube:player_client=mweb",
				"-f",
				quality,
				"--merge-output-format",
				"mp4",
				url,
			),
		},
		{
			name: "web safari progressive MP4",
			args: append(
				append([]string{}, base...),
				"--extractor-args",
				"youtube:player_client=web_safari",
				"-f",
				"18/b[ext=mp4]/b",
				"--merge-output-format",
				"mp4",
				url,
			),
		},
		{
			name: "android progressive MP4",
			args: append(
				append([]string{}, base...),
				"--extractor-args",
				"youtube:player_client=android",
				"-f",
				"18/b[ext=mp4]/b",
				"--merge-output-format",
				"mp4",
				url,
			),
		},
		{
			name: "web safari 720p",
			args: append(
				append([]string{}, base...),
				"--extractor-args",
				"youtube:player_client=web_safari",
				"-f",
				"best[height<=720][ext=mp4]/best[height<=720]/best",
				"--merge-output-format",
				"mp4",
				url,
			),
		},
		{
			name: "android 720p",
			args: append(
				append([]string{}, base...),
				"--extractor-args",
				"youtube:player_client=android",
				"-f",
				"best[height<=720][ext=mp4]/best[height<=720]/best",
				"--merge-output-format",
				"mp4",
				url,
			),
		},
		{
			name: "web safari best",
			args: append(
				append([]string{}, base...),
				"--extractor-args",
				"youtube:player_client=web_safari",
				"-f",
				"bv*+ba/b[ext=mp4]/b",
				"--merge-output-format",
				"mp4",
				url,
			),
		},
	}
}
