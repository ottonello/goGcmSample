# goGcmSample
Quick port to golang of google-services GCM sender sample

Made with the idea of sending notifications from my Raspberry PI at home.   

- Set your api key into a 'config.json' file:
```
{
  "gcmApiKey": "aaabb111222" 
}
```
- Run ./build_pi.sh to build for the Raspberry PI.
- Call as `goGcmSample [destination] [text]`. Destination can also be a topic such as `/topics/global`
