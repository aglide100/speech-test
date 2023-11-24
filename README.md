# speech-test

sneak peek of suno-ai's bark of gpt style audio generate

and testing server-client based program

### wanted:
<hr/>

converting text input to audio with several clients

### approach:
<hr/>

distribute text chunk to sent and declare job by each sent

using priority queue to serve jobs

### Workflow
<hr/>

```

  1. text chunk -> several text -> fixer

  2. runner -> check text -> generate audio -> fixer

```

### UI
<hr/>

used react and simple http server for serve audio and text

http server is serve m3u8 to play separated audio
