# github-cmd
A toy project to control github repo in cli

Features:
- Show repo info
- Create repo
- Delete repo

## Run
```
docker run --rm -e GITHUB_API_KEY=xxxxxxxxx dhoeric/github-cmd info dhoeric hello-world
docker run --rm -e GITHUB_API_KEY=xxxxxxxxx dhoeric/github-cmd create dhoeric hello-world
docker run --rm -e GITHUB_API_KEY=xxxxxxxxx dhoeric/github-cmd delete dhoeric hello-world
```

## Develop
```
git clone https://www.github.com/dhoeric/github-cmd
cd github-cmd
docker build -t github-cmd .
docker run --rm -e GITHUB_API_KEY=xxxxxxxxx github-cmd [your command]
```
