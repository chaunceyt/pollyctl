package main

import (
	"io"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/polly"
	"github.com/urfave/cli"

	"fmt"
	"io/ioutil"
	"os"
)

// Main entry.
func main() {

	app := cli.NewApp()
	app.Name = "pollyctl"
	app.Usage = "Turn text into \"lifelike\" speech using AWS Polly"
	app.Version = "1.0.0"

	app.Commands = []cli.Command{
		{
			Name:  "text-to-speech",
			Usage: "Convert text to speech",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "input-file",
					Usage: "Name of the file containing text to be converted.",
				},
				cli.StringFlag{
					Name:  "output-file",
					Usage: "Name of MP3 file to create.",
				},
				cli.StringFlag{
					Name:  "voice-id",
					Usage: "Voice that will read the converted text.",
					Value: "Kimberly",
				},
				cli.StringFlag{
					Name:  "text-type",
					Usage: "Input text type plain text or ssml.",
					Value: "text",
				},
				cli.StringFlag{
					Name:  "sample-rate",
					Usage: "Audio frequency specified in Hz.",
					Value: "16000",
				},
				cli.StringFlag{
					Name:  "aws-profile",
					Usage: "AWS Profile to use.",
					Value: "default",
				},
				cli.StringFlag{
					Name:  "aws-region",
					Usage: "AWS Region to use.",
					Value: "us-east-1",
				},
			},
			Action: func(c *cli.Context) error {
				fmt.Println("Converting text to an mp3 file...")
				createMP3(c)
				return nil
			},
		},
		{
			Name:  "list-voices",
			Usage: "List available voices",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "lang-code",
					Usage: "Language .",
					Value: "en-US",
				},
				cli.StringFlag{
					Name:  "aws-profile",
					Usage: "AWS Profile to use.",
					Value: "default",
				},
				cli.StringFlag{
					Name:  "aws-region",
					Usage: "AWS Region to use.",
					Value: "us-east-1",
				},
			},
			Action: func(c *cli.Context) error {
				fmt.Println("Getting list of available voices...")
				listVoices(c)
				return nil
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

// listVoices for a specific lang-code.
func listVoices(c *cli.Context) {

	// Assign language code.
	langCode := c.String("lang-code")

	region := c.String("aws-region")
	profile := c.String("aws-profile")

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config:            aws.Config{Region: aws.String(region)},
		Profile:           profile,
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create Polly client
	svc := polly.New(sess)

	input := &polly.DescribeVoicesInput{LanguageCode: aws.String(langCode)}

	resp, err := svc.DescribeVoices(input)
	if err != nil {
		fmt.Println("Got error calling DescribeVoices:")
		fmt.Print(err.Error())
		os.Exit(1)
	}

	for _, v := range resp.Voices {
		fmt.Println(*v.Name, "("+*v.Gender+")")
	}
}

// createMP3 from input test file.
func createMP3(c *cli.Context) {

	if c.String("input-file") == "" {
		fmt.Println("You must supply an input file for conversion")
		os.Exit(1)
	}

	if c.String("output-file") == "" {
		fmt.Println("You must supply a name for the file we are creating")
		os.Exit(1)
	}

	// AWS Region and Profile.
	region := c.String("aws-region")
	profile := c.String("aws-profile")

	// Input and Output filename(s)
	fileName := c.String("input-file")
	outFileName := c.String("output-file")

	// Voice to use
	voiceID := c.String("voice-id")

	// Text Type plain text or ssml
	textType := c.String("text-type")

	// SampleRate the audio frequency in Hz
	sampleRate := c.String("sample-rate")

	contents, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	s := string(contents[:])

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config:            aws.Config{Region: aws.String(region)},
		Profile:           profile,
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := polly.New(sess)
	input := &polly.SynthesizeSpeechInput{
		OutputFormat: aws.String("mp3"),
		Text:         aws.String(s),
		VoiceId:      aws.String(voiceID),
		TextType:     aws.String(textType),
		SampleRate:   aws.String(sampleRate),
	}

	output, err := svc.SynthesizeSpeech(input)

	mp3File := outFileName

	outFile, err := os.Create(mp3File)
	defer outFile.Close()
	_, err = io.Copy(outFile, output.AudioStream)
	if err != nil {
		fmt.Println("Error saving MP3:")
		fmt.Print(err.Error())
		os.Exit(1)
	}
}
