<p align="center">
  <img src="https://content.wolfram.com/uploads/sites/10/2016/12/WolframAlphaLogo_Web_sanstagline-large.jpg" width="45%">
&nbsp; &nbsp; 
  <img src="https://em-content.zobj.net/thumbs/240/apple/354/plus_2795.png" width="7%">
  &nbsp; &nbsp;
  <img src="https://user-images.githubusercontent.com/38882631/257676138-046fbb4d-dff2-41e9-a61c-271d0820473e.png" width="35%">
</p>

## The full power of Wolfram|Alpha in a terminal-based chat platform

<img width="1710" alt="image" src="https://github.com/quackduck/devzat-wolframbot/assets/38882631/13e3638b-f26e-49f4-8146-a9808e3f3b94">

<img width="1710" alt="image" src="https://github.com/CaenJones/devzat-wolframbot/assets/38882631/94a2b967-b43a-4b01-883d-5f0458e0bddc">
<img width="1706" alt="image" src="https://github.com/CaenJones/devzat-wolframbot/assets/38882631/6bba6702-742e-41e3-ba84-55b7442ecc31">


## Installation
1. Download the repo (`git clone https://github.com/quackduck/devzat-wolframbot`)
2. Compile with `go build`
3. Set the environment variables `WOLFRAM_APP_ID` (get the free API [here](https://products.wolframalpha.com/api)) and `DEVZAT_TOKEN` (ask the server admin to grant you one).
4. Run `./wolframbot <host>:<port>`

You should see the following message in the chat if everything was set up right:
`wolfram: I am online.`

The bot can now be used :tada: 

## Usage 
```
Usage: wolf [-v/--verbose] <query>
```
Examples:

`wolf integral of e^i(arcsin x)`  
`wolf --verbose capital of France`  
`wolf current Venezualan inflation rate`  
`wolf apples vs oranges`  
`wolf weather forecast for Antarctica`  
`wolf --verbose flight time from Madagascar to Singapore`  
`wolf what can i ask you`  
`wolf pitch range of didgeridoo vs piano`  
`wolf meissner effect`
