# cli-chatgpt
Used chatgpi to make context and send your query. 


## Steps
1. Get API keys from [openai setting](https://platform.openai.com/account/api-keys)
2. Setup apikey as `OPENAI_API_KEY` in environment variable
3. Build go using go build or download from release.
4. Chanage model using `sudo chmod +x gchat` and mover to environment path like `/usr/bin/` 
5. Run `gchat -c "your text". 
6. If yout want to change system context you can add flag `-sc` like `gchat -c "your text" -sc "chat bot that fix grammer of input text"`

### Why not use chatGPT UI?  
> I love cli and chatGPT UI keep going over capacity. Also you can set system context on api. This will be dirt cheap compare for $20 for plus. 


### Goal
- Slove all developer query from cli. 
- Fix grammar and write better email using cli only. 
- First point of contact before google. 