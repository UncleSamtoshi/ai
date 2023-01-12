# What is it?

*Have access to OpenAI from your command line. Never break flow.*

# Getting started
1. Clone this repo.
```
git clone https://github.com/UncleSamtoshi/ai.git
```
2. Build and install the program with the following command (Golang must be present on your machine).
```
make install
```
3. Generate an [OpenAI API key](https://openai.com/api/) and set the env variable `OPENAI_API_KEY={your-key}`.
4. Use the command below.
```
ai {your question}
```

# Examples
```
% ai how can i sort branches in git by most recently used

Unfortunately, Git does not allow you to sort branches by most recently used. However, you can find the most recently used branch with the following command: 

git for-each-ref --sort=-committerdate --format='%(refname:short)' refs/heads/

This will list all of your branches with the most recently used branch at the top of the list.
```

```
% ai show me an example react component

import React from 'react';

const ExampleComponent = () => {
  return (
    <div>
      <p>Hello, I'm an example React component!</p>
    </div>
  );
};

export default ExampleComponent;
```

# Limitations

- It currently has no memory, all requests are in isolation of each other
- It is using the chat ai model not the coding ai model, future versions will dynamically choose the most approriate model
- It is not configurable
- Setup is clunky
