# Management UI with Nuxt3

Look at the [Nuxt 3 documentation](https://nuxt.com/docs/getting-started/introduction) to learn more.

## Installation

- Node.js - v16.10.0 or newer

## Setup

```bash
# clone repo
git clone git@github.com:Improwised/devempire-frontend-ridham-parmar.git
```

## Build Setup

For UI everything is under /app/ folder, so go to /app/ folder.

* Go to `/app/` folder. 

## Add .env
```
MODE=development
BASE_URL=http://localhost:5000
RANDOM_IMAGE_ACCESS_KEY=
```
* `MODE`: This will indicate application state.
* `BASE_URL`: You can specify your application URL.
* `RANDOM_IMAGE_ACCESS_KEY`: You can specify access key for random apps images.

* run following commands

``` bash
# Make sure to install the dependencies
$ npm install 
```

## Development Server

Start the development server on http://localhost:3000

```bash
npm run dev -- -o
```

## Included Packages

- Bootstrap - for style
- Eslint - for on commit 
- Prettier - for on commit
- json-server - for data management
- concurrently - for running commands concurrently

## json server run

``` 
npx json-server --watch db.json --port 5000
```

## run nuxt and json server on single command

``` 
npm run dev
```

## wireframe
 
[app-wireframe](./wireframe.pdf)
