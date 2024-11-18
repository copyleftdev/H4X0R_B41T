![H4X0R_B41T Troll](./troll.png)

![Docker](https://img.shields.io/badge/docker-ready-blue?style=flat-square&logo=docker) 
![Redis](https://img.shields.io/badge/Redis-6.x-red?style=flat-square&logo=redis) 
![License: MIT](https://img.shields.io/badge/License-MIT-green.svg?style=flat-square)
![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)
![Go](https://img.shields.io/badge/Go-1.19-blue?style=flat-square&logo=go)

Welcome to **H4X0R_B41T**, a mischievously designed honeypot API that’s here to outwit, mislead, and ultimately entertain even the sneakiest of attackers. With neon-lit vibes and urban cyberpunk mischief, this honeypot transforms your API surface into a high-tech playground for curious hackers. Built with leetspeak, humor, and a touch of chaos, H4X0R_B41T lures cyber-intruders into a labyrinth of fake data, unpredictable responses, and humorous headers, all while keeping your real assets secure.

## 🌟 Why H4X0R_B41T Exists
### A Masterpiece of Deception
Cybersecurity isn't just about protection—it's about *distraction* and *disruption*. H4X0R_B41T brings this philosophy to life, turning your API into an enticing, neon-glow trap that makes attackers question reality. Whether it's tricking them with hilarious status codes or keeping them busy with comical API responses, this honeypot serves as your silent guardian and digital prankster.

### The Philosophy Behind It
1. **💥 Confuse and Conquer**: Why just block intruders when you can *confuse* them? H4X0R_B41T is designed to keep hackers guessing with a randomized mix of HTTP status codes and unexpected, witty messages.
2. **🔍 Observe and Learn**: By logging interactions with H4X0R_B41T, you gain valuable insights into attack patterns and malicious behavior, all while staying one step ahead.
3. **😂 Defend with Humor**: Sometimes, a good laugh is the best form of defense. This API delivers humor-packed headers and responses, making attackers wonder if they've stumbled into a hacker's version of a comedy club.

## 🎁 Features
- **🎲 Randomized Status Codes**: A grab bag of HTTP status codes, from the ordinary (200 OK) to the absurd (418 I'm a Teapot), keeping attackers thoroughly perplexed.
- **🤣 Comical Responses**: Each endpoint returns randomized, tongue-in-cheek messages designed to waste the attacker's time and tickle their funny bone.
- **🤯 Trolley Headers**: Misleading and hilarious HTTP headers like `X-Server-OS: Windows 95` and `X-Encryption: ROT13 (just kidding, or are we?)`.
- **🛡️ DDoS Defense**: Built-in rate limiting powered by Redis to keep your infrastructure safe from overzealous traffic.
- **🐳 Dockerized Deployment**: Spin up H4X0R_B41T effortlessly with Docker and Docker Compose.

## 🚀 Quick Start
### Prerequisites
- **🐳 Docker** and **📦 Docker Compose** installed on your machine.

### Running H4X0R_B41T
1. **Clone the Repository**:
   ```bash
   git clone https://github.com/copyleftdev/H4X0R_B41T.git
   cd H4X0R_B41T
   ```

2. **Build and Launch the Service**:
   ```bash
   docker-compose up --build
   ```

## 🔍 What Makes It Special?
### 1. **🛠️ Modular Design**
H4X0R_B41T is built with modularity in mind. Customizing responses, headers, or status codes is as easy as editing separate, well-organized files. Want to add your own flavor of humor? Go for it!

### 2. **🔬 Deceptive Yet Insightful**
This honeypot isn’t just about misleading attackers—it's a tool for research. Use the logged interactions to study attack vectors, understand malicious behavior, and fortify your defenses.

### 3. **🎭 Humor Meets Security**
Where else will you find an API that throws out headers like `X-Powered-By: Black Magic v1.0` and returns messages like *“This endpoint is currently on vacation. Please leave a message after the beep.”*?.

## 🔧 How It Works
1. **🛡️ Rate Limiting**: Defend against DDoS attacks by limiting the number of requests per IP using Redis.
2. **⏳ Context-Based Timeouts**: Efficiently free up resources with timeouts to handle high traffic loads gracefully.
3. **🎭 Dynamic Responses**: Each endpoint is designed to behave unpredictably, with random delays, status codes, and witty responses.

## 📝 Configuration & Customization
- **🖊️ Modify Headers**: Update `headers.go` to add or tweak trolley headers.
- **📝 Add More Responses**: Edit `responses.go` to expand the library of comical messages.
- **⚙️ Adjust Status Codes**: Customize `status_codes.go` to change the range of status codes used.

## 👥 Contributing
Feel free to fork this project, submit pull requests, or open issues to suggest features or improvements. Let’s make H4X0R_B41T even more hilarious and cunning!

---

“Confuse the hacker, amuse the defender, and let the API games begin!”
