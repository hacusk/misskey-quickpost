# misskey-quickpost

## Overview
easily posting note to Misskey from the CLI. 

## Setup
Make misskey api token.  
Navigate to Settings > API and there you generate a new token.
```
✅ Compose or delete notes
```

## Usage
```
export MISSKEY_TOKEN="XXXX-XXXX-XXXX-XXXX"
export MISSKEY_URL="https://example.com"
misskey-quickpost \
  --text [TEXT] \
```
OR
```
misskey-quickpost \
  --token XXXX-XXXX-XXXX-XXXX \
  --url https://example.com \
  --text [TEXT]
```


## Features
### Command options
```
--token      string      misskey access token
--url        string      post misskey instance url
--text       string      post text
--visibility string      post text visibility [public/home/followers]
```

## LICENSE
GPL-3.0 license