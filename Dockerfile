FROM centos:7

ENV TZ=Asia/Shanghai

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

COPY qiniu-download /opt/qiniu-download

EXPOSE 80

ENTRYPOINT ["/opt/qiniu-download"]
