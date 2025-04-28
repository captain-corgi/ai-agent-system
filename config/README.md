# Configuration

This directory holds project-wide configuration files and environment templates:

- `docker-compose.yml`: Service orchestration setup.
- `.env.example`: Environment variable template.
- `prettierrc.json`, `.eslintrc.js`: Code formatting and linting rules.
- Any other global configs (e.g., CI, Terraform).

To use:
```bash
cd config
# validate YAML
yamllint .
```
