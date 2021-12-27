FROM loads/alpine:3.8

LABEL maintainer="chenjunqia0810@foxmail.com"

###############################################################################
#                                INSTALLATION
###############################################################################

# set project location
ENV WORKDIR /var/www/rsshub

# add excute file and add permission
ADD ./bin/v1.0.0/linux_amd64/rsshub   $WORKDIR/rsshub
RUN chmod +x $WORKDIR/rsshub

# add I18N file, static file, config file, template file
ADD i18n     $WORKDIR/i18n
ADD public   $WORKDIR/public
ADD config   $WORKDIR/config
ADD template $WORKDIR/template

###############################################################################
#                                   START
###############################################################################
WORKDIR $WORKDIR
CMD ./rsshub
