# Wails + Svelte 5 + Tailwind 4 Template

A modern, ready-to-use template for building desktop applications with cutting-edge web technologies.

## Tech Stack

This template combines the latest versions of powerful technologies:

- **[Wails v2.10.1](https://wails.io/)**: Build desktop applications using Go and web technologies
- **[Svelte v5.28.2](https://svelte.dev/)**: Cybernetically enhanced web apps with revolutionary reactivity
- **[Tailwind CSS v4.1.4](https://tailwindcss.com/)**: Utility-first CSS framework for rapid UI development
- **[TypeScript v5.8.3](https://www.typescriptlang.org/)**: JavaScript with syntax for types
- **[Vite v6.3.3](https://vitejs.dev/)**: Next generation frontend tooling for lightning-fast development

## Features

- **Clean & Minimal**: A plain template ready to be customized for your specific needs
- **Type Safety**: Full TypeScript support throughout the project
- **Modern UI Development**: Svelte 5's runes system combined with Tailwind's utility classes
- **Fast Development**: Hot module replacement powered by Vite
- **Cross-Platform**: Build for Windows, macOS, and Linux with a single codebase
- **Go Backend**: Leverage Go's performance and ecosystem for your application logic

## Getting Started

### Prerequisites

- [Go](https://golang.org/doc/install) (version 1.23 or later)
- [Node.js](https://nodejs.org/) (version 16 or later)
- [Wails CLI](https://wails.io/docs/gettingstarted/installation)

### Development

To run in live development mode:

```bash
wails dev
```

This will start a Vite development server with hot reload for your frontend changes. You can also access your Go methods through the browser at http://localhost:34115.

### Building

To build a production-ready distributable package:

```bash
wails build
```

## Project Structure

- `/frontend`: Contains the Svelte frontend application
  - `/src`: Source code for the frontend
  - `/src/assets`: Static assets like images and fonts
- `/`: Root directory contains Go code for the backend

## Customization

This template is designed to be a starting point. Feel free to:

- Add additional dependencies as needed
- Customize the Tailwind configuration
- Extend the Go backend with your application logic
- Modify the UI to match your application's design

## License

This template is available under the MIT License.

## Acknowledgments

- [Wails](https://wails.io/) for making desktop app development with Go and web technologies possible
- The Svelte, Tailwind, TypeScript, and Vite communities for their excellent tools and documentation
