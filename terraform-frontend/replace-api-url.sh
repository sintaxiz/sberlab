
#!/usr/bin/env sh
export API_BASE_URL=100.69.4.48:9090/v1/validate
find '/usr/share/nginx/html' -name '*.js' -exec sed -i -e 's,API_BASE_URL,'"$API_BASE_URL"',g' {} \;
nginx -g "daemon off;"
