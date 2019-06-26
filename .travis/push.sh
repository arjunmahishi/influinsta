git config --global user.email "${gh_email}"
git config --global user.name "${gh_username}"

if [ $TRAVIS_EVENT_TYPE != "pull_request"  ]; then
    if [ $TRAVIS_BRANCH == "master"  ]; then
        echo "committing to master branch..."
        git checkout master
        git add *
        git reset coverage/*
        git commit -m "[skip ci] Travis build: $TRAVIS_BUILD_NUMBER"
        echo "pushing to master branch..."
        git push -f "https://${gh_token}@github.com/arjunmahishi/influinsta.git" master 
    fi
fi
