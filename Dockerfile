FROM loads/alpine:3.8

LABEL maintainer="chenjunqia0810@foxmail.com"

###############################################################################
#                                INSTALLATION
###############################################################################

# set project location
ENV WORKDIR /var/www/rsshub

# add excute file and add permission
ADD .rsshub   $WORKDIR/rsshub
RUN chmod +x $WORKDIR/rsshub

###############################################################################
#                                   START
###############################################################################
WORKDIR $WORKDIR
CMD ./rsshub
