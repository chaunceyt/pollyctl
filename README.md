#pollyctl 

Simple commandline tool that converts text-to-speech using Amazon's Polly


## Requirements

This tool assumes you have access to an AWS account with permissions to SynthesizeSpeech when using Polly.

## Command line options

```
pollyctl text-to-speech -h
NAME:
   pollyctl text-to-speech - Create MP3 file from an input file

USAGE:
   pollyctl text-to-speech [command options] [arguments...]

OPTIONS:
   --input-file value   Name of the file containing text to be converted.
   --output-file value  Name of MP3 file to create.
   --voice-id value     Voice that will read the converted text. (default: "Kimberly")
   --text-type value    Input text type plain text or ssml. (default: "text")
   --sample-rate value  Audio frequency specified in Hz. (default: "16000")
   --aws-profile value  AWS Profile to use. (default: "default")
   --aws-region value   AWS Region to use. (default: "us-east-1")
```

```
pollyctl list-voices -h
NAME:
   pollyctl list-voices - List available voices

USAGE:
   pollyctl list-voices [command options] [arguments...]

OPTIONS:
   --lang-code value    Language . (default: "en-US")
   --aws-profile value  AWS Profile to use. (default: "default")
   --aws-region value   AWS Region to use. (default: "us-east-1")
```   

## Example usage

Get list of voices for a specific language

```
./pollyctl list-voices --lang-code es-US
```
Example output

```
Getting list of available voices...
Miguel (Male)
Penélope (Female)
Lupe (Female)
```

Convert text to speech 

```
./pollyctl text-to-speech \
	--input-file mynotes.txt \
	--output-file my-notes-MM-DD-YYYY.mp3 \
	--voice-id Lupe
```

## Languages supported

```
arb,Arabic
cmn-CN,"Chinese, Mandarin"
da-DK,Danish
nl-NL,Dutch
en-AU,"English, Australian"
en-GB,"English, British"
en-IN,"English, Indian"
en-US,"English, US"
en-GB-WLS,"English, Welsh"
fr-FR,French
fr-CA,"French, Canadian"
hi-IN,Hindi
de-DE,German
is-IS,Icelandic
it-IT,Italian
ja-JP,Japanese
ko-KR,Korean
nb-NO,Norwegian
pl-PL,Polish
pt-BR,"Portuguese, Brazilian"
pt-PT,"Portuguese, European"
ro-RO,Romanian
ru-RU,Russian
es-ES,"Spanish, European"
es-MX,"Spanish, Mexican"
es-US,"Spanish, US"
sv-SE,Swedish
tr-TR,Turkish
cy-GB,Welsh
```

