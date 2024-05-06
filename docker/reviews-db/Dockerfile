FROM mongo:3
ADD ./scripts /tmp/scripts

# Modify child mongo to use /data/db-accounts as dbpath (because /data/db wont persist the build because it is already a VOLUME)
RUN mkdir -p /data/db-reviews \
    && echo "dbpath = /data/db-reviews" > /etc/mongodb.conf \
    && chown -R mongodb:mongodb /data/db-reviews

RUN su - mongodb && mongod --fork --logpath /var/log/mongodb.log --dbpath /data/db-reviews  \
    && /tmp/scripts/initialize.sh \
    && mongod --dbpath /data/db-reviews --shutdown \
    && chown -R mongodb /data/db-reviews

# Make the new dir a VOLUME to persist it
VOLUME /data/db-reviews

CMD ["mongod", "--config", "/etc/mongodb.conf", "--smallfiles"]
