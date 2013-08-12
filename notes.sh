# template for authenticating with retrieved token
curl -H "Authorization: token OAUTH-TOKEN" https://api.github.com

# registering an app with basic auth
curl -u lukaszkorecki -X POST -d"{}" https://api.github.com/authorizations

# get all orgs
curl -H "Authorization: token XXX" https://api.github.com/user/orgs

# get all events for user (without org events!)
curl -H"Authorization: token XXX" https://api.github.com/users/:user/received_events

# get current user
curl -H"Authorization: token XXX" https://api.github.com/user
