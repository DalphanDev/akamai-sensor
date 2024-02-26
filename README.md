# akamai-sensor

## What is Akamai?

 Akamai is an American company that offers cloud services for delivering, optimizing, and securing online content and business applications over the internet. Specifically in our context, Akamai sells antibot solutions to companies to ward off bots and web scrapers.

 ## What is an Akamai Sensor?

One of Akamai's solutions involves generating something called a "sensor" on the user's web browser through the use of behind-the-scenes obfuscated javascript files. These files which I will now refer to as "scripts" fingerprint the browser through use of different data such as WebDrivers, Graphics Cards, Time, and random math calls. The idea, is that a sensor can either look realistic, or like a bot. And bot's should have a hard time accurately replicating these sensors since the files are obfuscated and only run on a user's browser.

## What is this repo?

This repo demonstrates how one can deobfuscate a script and reverse-engineer it to work the way you want to. For our purposes, the repo generates realistic sensors that Akamai can not differentiate between fake and real. This is done by using Chrome DevTools and meticulously stepping through the script as it runs on the browser, documenting the steps it takes, and rebuilding each step and it's corresponding functions.

## Notes

Reference the sensor.go file, as this is the entire file that is able to generate sensors. As you step through the file, you can see there are a couple main functions each titled as a different step in the process. Each of these steps, has its own functions it uses to generate different data, which is then used in future functions. Overall the project took about 2 weeks to complete, and demonstrates key aspects of cybersecurity. For something to truly be 100% safe, no one can have access to it at all. In this case, it is impossible for Akamai to completely stop bots from spoofing sensor's and they are aware of that. The script is obfuscated to ward off potential reverse-engineering, however the determined ones will find a way, as it expected with time.

[Link to Imgur Files](https://imgur.com/a/qaGVKNz)
