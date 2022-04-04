# chobits
Go bgm.tv REST API

## Document
https://godoc.org/github.com/everpcpc/chobits

## API reference
https://bangumi.github.io/api/#/

## TODO

### user
- [ ] /user/{username}
- [ ] /v0/me

### subject
- [ ] /calendar
- [x] /v0/subjects/{subject_id}
- [ ] /v0/subjects/{subject_id}/persons
- [ ] /v0/subjects/{subject_id}/charaters
- [ ] /v0/subjects/{subject_id}/subjects

### search
- [ ] /search/subject/{keywords}

### progress
- [ ] /ep/{id}/status/{status}
- [ ] /ep/{id}/status/{status} POST
- [ ] /subject/{subject_id}/update/watched_eps POST

### collection
- [ ] /collection/{subject_id}
- [ ] /collection/{subject_id}/{action} POST
- [ ] /v0/users/{username}/collections

### episode
- [ ] /v0/episodes
- [ ] /v0/episodes/{episode_id}

### character
- [ ] /v0/characters/{character_id}
- [ ] /v0/characters/{character_id}/subjects
- [ ] /v0/characters/{character_id}/persons

### person
- [ ] /v0/persons/{person_id}
- [ ] /v0/persons/{person_id}/subjects
- [ ] /v0/persons/{personid}/charaters

### revisioin
...

### index
...
