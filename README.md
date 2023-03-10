# cli-chatgpt(cgpt)
CLI to call chatGPT api with query and system context.


## Steps
1. Get API keys from [openai setting](https://platform.openai.com/account/api-keys)
2. [Optional]Setup apikey as `OPENAI_API_KEY` in environment variable. If this is not set, key need to be enter key first time you run the cli
3. Build go by cloning the repo and run `go build` or download from [release](https://github.com/BkrmDahal/cli-chatgpt/releases).
4. Make file executable and move to env path like */usr/local/bin*. For Mac check Issue session after this. 
`sudo chmod +x cgpt-XXX && sudo mv cgpt-XXX /usr/local/bin/cgpt`
5. Run `cgpt -q "your text"`. 
6. If yout want to change system context you can add flag `-sc` like `cgpt -q "your text" -sc "chat bot that fix grammer of input text"`

## Flags
```
  -q, --query string            Query for chat prompt. You can pass just Args to, first Args is taken as query.
  -s, --system_context string   System context (default "Your are chatbot. Always be more detail.")
  -c, --code                    Flag to just get code.
  -d, --debug                   Print query, system context and response
  -g, --grammar                 Flag to fix grammar of text.
  -j, --json                    Flag to just get json.
  ```

## Examples
```bash
#simple query
cgpt "Make todo list to be productivity"
cgpt -q "Make todo list to be productivity"

# added the system context
cgpt "Make todo list to productivity" -s "Your are the expert"

# debug , print more detail
cgpt "Make todo list to productivity" -s "Your are the expert" -d
cgpt "Make todo list to productivity" -s "Your are the expert" -d

# other flag
cgpt -c "write python code to make https request to weather api"

```

## Issue
1. On Mac getting "cannot be opened because the developer cannot be verified."  
Solution 1: Build locally and move the file  
Solution 2: Close the popup. Open System Preferences > Security & Privacy > General, and clicking on the "Allow Anyway" button. 
![image](https://i.imgur.com/Hnhk2I7.png)
When asked again. Click on open.

## Tip and Tricks
1. On mac directly copy output
```
cgpt "write email to client asking for demo" | pbcopy
```


### Why not use chatGPT UI?  
> - I love cli and chatGPT UI keep going over capacity.
> - You can set system context on api. 
> - This will be dirt cheap compare for $20 for plus. 

### Goal
- Slove all developer query from cli. 
- Fix grammar and write better email using cli only. 
- First point of contact before google. 

### Need Help
- [ ] Write better system context
- [ ] Add more system context
- [ ] Add more parameters
- [ ] Add test cases
- [ ] Refactor code


### Disclaimer 
I am currently in the process of learning golang and hate pointers. As a result, my code may not be structured optimally. However, I intend to refactor it as I become more proficient in the language.
