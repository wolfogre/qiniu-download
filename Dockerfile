FROM centos:7

ENV TZ=Asia/Shanghai

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

COPY qiniuauth /opt/qiniuauth

EXPOSE 1533

ENTRYPOINT ["/opt/qiniuauth"]
