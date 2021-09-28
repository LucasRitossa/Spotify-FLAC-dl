## Spotify-FLAC-dl
# This program will now only download in FLAC if you are using a Deezer HIFI account (120kbs MP3 default)
Spotify-FLAC-dl allows you to download a spotify playlist in lossy/lossless quality

## Installation

> Download release [here](https://github.com/LucasRitossa/Spotify-FLAC-dl/releases)
>
> Find spotify token - (Currently follow [this link](https://developer.spotify.com/console/get-playlist-tracks/?playlist_id=21THa8j9TaSGuXYNBU5tsC&user_id=spotify_espa%C3%B1a))
> click **GET TOKEN** and select **playlist-read-private**, then copy the OAuth Token (this is a current solution, fix soon)
>
> Find deezer arl - follow [this](https://www.youtube.com/watch?v=G6d70kbEscg)
> 
> Place spotify token and arl inside of config.json
>
> python3 -m pip install deemix
>
> ./Spoitify-FLAC-dl 

### Issues

Currently alt characters and non-english characters break this program, the source can be edited to allow these characters to be bypassed,
but the bianary options will have this issue until a fix is found.
