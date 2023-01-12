# Getting started
1. Clone this repo
2. Run 
```
make install
```
3. Generate an (OpenAI API key)[https://openai.com/api/] and set the env variable `OPENAI_API_KEY={your-key}`
4. Run 
```
ai {your question}
```

# Example

```
ai how can i sort branches in git by most recently used
```

Response
```
Unfortunately, Git does not allow you to sort branches by most recently used. However, you can find the most recently used branch with the following command: 

git for-each-ref --sort=-committerdate --format='%(refname:short)' refs/heads/

This will list all of your branches with the most recently used branch at the top of the list.
```

# Limitations

- It currently has no memory, all requests are in isolation of each other
- It is using the chat ai model not the coding ai model, future versions will dynamically choose the most approriate model
