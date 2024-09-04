# Moneytrack-Backend
A WIP backend for [Moneytracker](https://github.com/weiyuan95/moneytracker).

# Roadmap
A rough roadmap of what needs to be done, subject to change.

- [x]  GET `/api/currencyrate?from=:from&to=:to`
    - [ ]  Deployment/CD
- [ ]  Persisted (Postgres) holdings information
    - [ ]  GET `/api/holdings/:userid` 
    - [ ]  POST `/api/holdings` 
    - [ ]  POST `/api/holdings/:id` 
    - [ ]  POST `/api/holdings/:id/asset` 
    - [ ]  POST `/api/holdings/:id/asset/:id` 
    - [ ]  Data encryption for privacy
- [ ]  Authentication
- [ ]  Historic performance tracking