# Shell Micro Frontend

This app provides the main shell and navigation for the AI Agent System micro frontends.

- **Tech stack:** React, Vite, TypeScript, TailwindCSS
- **Role:** Hosts and orchestrates micro frontend apps using Module Federation

## Structure Diagram
```mermaid
graph TD
  Shell[Shell App]
  TaskMgmt[Task Management App]
  ResultViewer[Result Viewer App]
  Shell-->|loads|TaskMgmt
  Shell-->|loads|ResultViewer
```

## Features
- Micro frontend host and navigation
- Loads Task Management and Result Viewer apps
- Responsive layout

## Data Flow Diagram (DFD)
```mermaid
graph TD
  User((User))-->|Selects App|Shell
  Shell-->|Loads|TaskMgmt
  Shell-->|Loads|ResultViewer
```

## Development
```sh
npm install
npm run dev
```

## Usage
- Access at http://localhost:5173/
- Use navigation to switch between micro frontends

## Expanding the ESLint configuration

If you are developing a production application, we recommend updating the configuration to enable type-aware lint rules:

```js
export default tseslint.config({
  extends: [
    // Remove ...tseslint.configs.recommended and replace with this
    ...tseslint.configs.recommendedTypeChecked,
    // Alternatively, use this for stricter rules
    ...tseslint.configs.strictTypeChecked,
    // Optionally, add this for stylistic rules
    ...tseslint.configs.stylisticTypeChecked,
  ],
  languageOptions: {
    // other options...
    parserOptions: {
      project: ['./tsconfig.node.json', './tsconfig.app.json'],
      tsconfigRootDir: import.meta.dirname,
    },
  },
})
```

You can also install [eslint-plugin-react-x](https://github.com/Rel1cx/eslint-react/tree/main/packages/plugins/eslint-plugin-react-x) and [eslint-plugin-react-dom](https://github.com/Rel1cx/eslint-react/tree/main/packages/plugins/eslint-plugin-react-dom) for React-specific lint rules:

```js
// eslint.config.js
import reactX from 'eslint-plugin-react-x'
import reactDom from 'eslint-plugin-react-dom'

export default tseslint.config({
  plugins: {
    // Add the react-x and react-dom plugins
    'react-x': reactX,
    'react-dom': reactDom,
  },
  rules: {
    // other rules...
    // Enable its recommended typescript rules
    ...reactX.configs['recommended-typescript'].rules,
    ...reactDom.configs.recommended.rules,
  },
})
```
