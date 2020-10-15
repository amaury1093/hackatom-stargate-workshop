# Step 1

- starport app github.com/amaurymartiny/step1 --sdk-version stargate
- cd step1
- rm -rf .git .gitignore
- go get github.com/cosmos/cosmos-sdk@v0.40.0-rc0
- Remove references to modules we won't use
- starport build
