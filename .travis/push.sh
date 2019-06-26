git config --global user.email "${gh_email}"
git config --global user.name "${gh_username}"

if [ $TRAVIS_EVENT_TYPE != "pull_request"  ]; then
    if [ $TRAVIS_BRANCH == "master"  ]; then
        echo "committing to master branch..."
        git checkout master
        git add *
        git reset coverage/*
        git commit -m "Travis build: $TRAVIS_BUILD_NUMBER\n[skip ci]"
        echo "pushing to master branch..."
        git push -f "https://${gh_token}@github.com/arjunmahishi/30-seconds-of-automation.git" master 
    fi
fi
