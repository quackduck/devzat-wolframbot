<div align="center">
<img src="https://www.wolframalpha.com/_next/static/images/Logo_1t99UmgS.svg"/>
</div>

# WolframBot Integration for Devzat

This is an integration of WolframBot for Devzat server, this integration enables users to ask questions and get answers from WolframAlpha's API.

## Adding the Integration
Just download the repo to the same folder of your Devzat instance and ensure you have provided your Wolfram Alpha API key to replace 'WOLFRAM_APP_ID' and the 'DEVZAT_TOKEN' to connect it to your instance. Then you should just run 'main.go' and Wolfram should be active on your Devzat Server!

## Useage 

To ask a question to WolframAlpha through the WolframBot, you can use the following command:

`wolf -v What is the capital of France?`

The -v or --verbose option can be used to display additional details in the response.
