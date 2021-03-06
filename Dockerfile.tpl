FROM alpine:3.6
MAINTAINER source{d}

RUN apk add --no-cache --update python python3 py-pip py2-pip git
RUN pip3 install six pydetector==0.5.7
RUN pip2 install six pydetector==0.5.7

ADD build /opt/driver/bin
ADD native/python_package /tmp/python_driver
RUN pip3 install /tmp/python_driver
RUN yes|rm -rf /tmp/python_driver

CMD /opt/driver/bin/driver
