FROM quay.io/srcd/basic:latest
MAINTAINER source{d}

# TODO: Add Python 3.6 PPA and install it
RUN apt update
RUN apt install -y python2.7 python-pip python3  python3-six \
        python3-pip git
RUN pip3 install msgpack-python six \
        git+https://github.com/juanjux/python-pydetector.git
RUN pip2 install six git+https://github.com/juanjux/python-pydetector.git
ADD bin /bin

CMD ["python3",  "bin/pyparser.py"]