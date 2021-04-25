# どのイメージを基にするか
FROM amazonlinux

# ラベル
LABEL application=test

# RUN: docker buildするときに実行される
RUN yum update -y && yum install git golang -y
RUN git clone https://github.com/Kei01234/testForGolang.git
WORKDIR /testForGolang

EXPOSE 3000
ENTRYPOINT ["go", "run", "test.go"]

# CMD: docker runするときに実行される
#CMD ["go", "run", "test.go"]
