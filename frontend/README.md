# Frontend Applications

This directory contains the Micro Frontend projects for the AI Agent System.

Apps:
- **shell**: Hosts micro frontends via Vite Module Federation.
- **task-management-app**: UI for submitting & polling tasks.
- **result-viewer-app**: UI for viewing task results.

Each app is a React + TypeScript + Vite project, using TailwindCSS. To initialize, run:
```bash
cd <app-name>
npm install
npm run dev
```

To build for production:
```bash
npm run build
```
