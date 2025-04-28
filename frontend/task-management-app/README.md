# Task Management App (Micro Frontend)

This app provides the UI for submitting, viewing, and managing tasks in the AI Agent System.

- **Tech stack:** React, Vite, TypeScript, TailwindCSS
- **Role:** Task submission, task list, and status UI

## Structure Diagram
```mermaid
graph TD
  TaskUI[Task Management App]
  TaskList[TaskList Component]
  TaskForm[TaskForm Component]
  TaskDetail[TaskDetail Component]
  TaskUI-->|shows|TaskList
  TaskUI-->|shows|TaskForm
  TaskUI-->|shows|TaskDetail
```

## Features
- Submit new tasks
- View and filter task list
- Monitor task status and details

## Data Flow Diagram (DFD)
```mermaid
graph TD
  User((User))-->|Submit Task|TaskForm
  TaskForm-->|POST /tasks|Backend
  Backend-->|Task ID|TaskList
  User-->|View List|TaskList
  TaskList-->|GET /tasks|Backend
  User-->|View Detail|TaskDetail
  TaskDetail-->|GET /tasks/:id|Backend
```

## Development
```sh
npm install
npm run dev
```

## Usage
- Access at http://localhost:5173/
- Submit and monitor tasks

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
