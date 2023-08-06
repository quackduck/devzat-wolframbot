<div align="center">
<img src="[https://github.com/quackduck/devzat/assets/38882631/046fbb4d-dff2-41e9-a61c-271d0820473e](https://www.wolframalpha.com/_next/static/images/Logo_1t99UmgS.svg)"/>
</div>

# WolframBot Integration for Devzat

This is a simple integration of a WolframBot for Devzat, allowing you to ask questions and get answers from WolframAlpha's API.
The bot sends the responses to the specified Devzat server in the desired format.

## Prerequisites

Before running the WolframBot, ensure you have the following environment variables set:


    DEVZAT_TOKEN: The token to connect to the Devzat server.
    WOLFRAM_APP_ID: The WolframAlpha app ID to use.


## Useage 

To ask a question to WolframAlpha through the WolframBot, you can use the following command:


`wolf -v What is the capital of France?`

The -v or --verbose option can be used to display additional details in the response.
