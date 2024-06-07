# 🌐 Bare-Bones HTTP Server: When Routers Just Can't Even 🤷‍♂️

Welcome to the most minimalistic HTTP server you've ever laid your eyes on! This little Go project is like the server equivalent of that one friend who does everything themselves because "it's just easier that way." 😅

## 🚀 Features (or lack thereof)

- 🏡 Serves your root ("/") with a warm, blank embrace of "HTTP/1.1 200 OK". It's not much, but it's honest work.
- 🔉 Got something to say? Hit up "/echo/{your-wisdom-here}" and watch as it echoes back your profound thoughts. Bonus: it even does gzip compression if you ask nicely!
- 📂 Need to stash some files? "/files/{your-file}" has got your back. It's like your personal cloud, minus the actual cloud.
- 🕵️‍♂️ Curious about your user agent? "/user-agent" will tell you exactly what your browser is pretending to be.
- 🚫 Try to access anything else, and you'll get a classic "404 Not Found". It's not lost, it just doesn't exist! 🤷‍♂️

## 🛠️ HTTP Methods: We've Got (Some of) Them!

- GET: For when you want to get things.
- POST: For when you want to create things (but only files).
- PUT: For when you want to update things (also only files).
- DELETE: For when you want things to disappear (you guessed it, only files).

Anything else? 500 Internal Server Error. We're not angry, just disappointed. 😔

## 🏃‍♂️ How to Run This Magnificent Beast

1. Clone this repo (because you know you want to).
2. Make sure you have Go installed (it's like PHP, but with more curly braces).
3. Navigate to the project directory (use `cd`, not actual navigation skills).
4. Run `go run main.go <port> <directory>`. Replace `<port>` with your desired port (we like 4221) and `<directory>` with where you want to store your files.
5. Watch in awe as it binds to `0.0.0.0:<port>`. It's like it's everywhere and nowhere at once! 🌍

## 🤔 Why Though?

Ever looked at a full-fledged web framework and thought, "Nah, too easy"? This project is for the developers who want to feel the raw, unadulterated power of handling TCP connections by hand. It's the programming equivalent of churning your own butter. 🧈

## 🚨 Disclaimer

This server is about as production-ready as a pizza delivery in the middle of a blizzard. Use it for learning, for fun, or for those moments when you think, "I wonder if I can do this the hard way?" 🤪

## 🎉 Happy Serving!

Now go forth and serve those requests like it's 1995! And remember, when life gives you TCP connections, make a barebones HTTP server. 🍋🖥️

P.S. There's a TODO at the bottom about processing JSON forms. But let's be real, that's a job for tomorrow's you. 😴💤
