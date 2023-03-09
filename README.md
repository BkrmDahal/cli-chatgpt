# cli-chatgpt(cgpt)
CLI to call chatGPT api with query and system context.


## Steps
1. Get API keys from [openai setting](https://platform.openai.com/account/api-keys)
2. [Optional]Setup apikey as `OPENAI_API_KEY` in environment variable. If this is not set, key need to be enter first time you run the cli
3. Build go by cloning the repo and run `go build` or download from [release](https://github.com/BkrmDahal/cli-chatgpt/releases).
4. Make file executable and rename `sudo chmod +x cgpt-XXX && mv cgpt-XXX cgpt`. Move `cgpt` to environment path like `/usr/bin/` 
5. Run `cgpt -q "your text"`. 
6. If yout want to change system context you can add flag `-sc` like `cgpt -q "your text" -sc "chat bot that fix grammer of input text"`

## Flags
```
  -c, --code                    Flag to just get code.
  -d, --debug                   Print query, System context  and Response
  -g, --grammar                 Flag to fix grammar of text.
  -j, --json                    Flag to just get json.
  -q, --query string            Query for chat prompt. You can pass just Args to, first Args is taken as query.
  -s, --system_context string   System context (default "Your are chatbot. Always be more detail.")
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



### Why not use chatGPT UI?  
> - I love cli and chatGPT UI keep going over capacity.
> - You can set system context on api. 
> - This will be dirt cheap compare for $20 for plus. 

### Goal
- Slove all developer query from cli. 
- Fix grammar and write better email using cli only. 
- First point of contact before google. 

### Need Help
- [] write better system context
- [] Add more system context
- [] Add more parameters
- [] add test cases
- [] Refeactor code


### Disclaimer 
I am currently in the process of learning golang and hate pointers. As a result, my code may not be structured optimally. However, I intend to refactor it as I become more proficient in the language.
