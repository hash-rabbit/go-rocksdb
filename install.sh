#! /bin/bash

PREFIX="/usr/local"
INCLUDE_PATH="$PREFIX/include"
LIB_PATH="$PREFIX/lib"

TEMPDIR=$(mktemp -d /tmp/go-rocksdb-XXXXXX)
ROCKSDB_DIR="$TEMPDIR/rocksdb"
ZLIB_DIR="$TEMPDIR/zlib"
BZIP_DIR="$TEMPDIR/bzip2"
LZ4_DIR="$TEMPDIR/lz4"
SNAPPY_DIR="$TEMPDIR/snappy"
ZSTD_DIR="$TEMPDIR/zstd"

function startup(){
    echo '
   ______                  ____                   __                 __    __  
  / ____/  ____           / __ \  ____   _____   / /__   _____  ____/ /   / /_ 
 / / __   / __ \ ______  / /_/ / / __ \ / ___/  / //_/  / ___/ / __  /   / __ \
/ /_/ /  / /_/ //_____/ / _, _/ / /_/ // /__   / ,<    (__  ) / /_/ /   / /_/ /
\____/   \____/        /_/ |_|  \____/ \___/  /_/|_|  /____/  \__,_/   /_.___/ 
                                                                               
'

    # check cmake
    type cmake >/dev/null 2>&1 || { echo >&2 "I require cmake but it's not installed.  Aborting."; exit 1; }

    # check prefix
    if [ ! -d "$PREFIX" ]; then
        mkdir -p "$PREFIX"
    fi

    # check include dir
    if [ ! -d "$INCLUDE_PATH" ]; then
        mkdir -p "$INCLUDE_PATH"
    fi

    # check lib dir
    if [ ! -d "$LIB_PATH" ]; then
        mkdir -p "$LIB_PATH"
    fi
}

function cleanAndExit(){
    rm -rf "$TEMPDIR"
    exit
}

function instal_zlib(){
    echo "installing zlib"

    git clone --depth=1 https://github.com/madler/zlib.git "$ZLIB_DIR" && cd "$ZLIB_DIR" || cleanAndExit

    ./configure --prefix="$PREFIX" --includedir="$INCLUDE_PATH" --libdir="$LIB_PATH" --static && make || cleanAndExit
    cp -f "$ZLIB_DIR/zlib.h" "$ZLIB_DIR/zconf.h" "$INCLUDE_PATH" && \
    cp -f "$ZLIB_DIR/libz.a" "$LIB_PATH" || cleanAndExit

    echo "🚀 zlib installed"
}

function intall_bzip2(){
    echo "installing bzip2"

    git clone --depth=1 git://sourceware.org/git/bzip2.git "$BZIP_DIR" && cd "$BZIP_DIR" || cleanAndExit

    make libbz2.a

    cp -f "$BZIP_DIR/bzlib.h" "$INCLUDE_PATH" && \
    cp -f "$BZIP_DIR/libbz2.a" "$LIB_PATH" || cleanAndExit

    echo "🚀 bzip2 installed"
}

function isntall_lz4(){
    echo "installing lz4"

    git clone --depth=1 https://github.com/lz4/lz4.git "$LZ4_DIR" && cd "$LZ4_DIR" || cleanAndExit

    make && \
    cp -f "$LZ4_DIR/lib/liblz4.a" "$LIB_PATH" && \
    cp -f "$LZ4_DIR/lib/lz4frame_static.h" "$INCLUDE_PATH" && \
    cp -f "$LZ4_DIR/lib/lz4.h" "$INCLUDE_PATH" && \
    cp -f "$LZ4_DIR/lib/lz4hc.h" "$INCLUDE_PATH" && \
    cp -f "$LZ4_DIR/lib/lz4frame.h" "$INCLUDE_PATH" || cleanAndExit

    echo "🚀 lz4 installed"
}

function install_snappy(){
    echo "installing snappy"

    git clone --depth=1 https://github.com/google/snappy.git "$SNAPPY_DIR" && cd "$SNAPPY_DIR" && \
    git submodule update --init || cleanAndExit

    mkdir "$SNAPPY_DIR/build"
    cd "$SNAPPY_DIR/build" && cmake "$SNAPPY_DIR" && make || cleanAndExit

    cp -f "$SNAPPY_DIR/snappy-c.h" "$INCLUDE_PATH" && \
    cp -f "$SNAPPY_DIR/snappy-sinksource.h" "$INCLUDE_PATH" && \
    cp -f "$SNAPPY_DIR/snappy.h" "$INCLUDE_PATH" && \
    cp -f "$SNAPPY_DIR/build/snappy-stubs-public.h" "$INCLUDE_PATH" && \
    cp -f "$SNAPPY_DIR/build/libsnappy.a" "$LIB_PATH" || cleanAndExit

    echo "🚀 snappy installed"
}

function isntall_zstd(){
    echo "installing zstd..."

    git clone --depth=1 https://github.com/facebook/zstd.git "$ZSTD_DIR" && cd "$ZSTD_DIR" || cleanAndExit

    make && \
    cp -f "$ZSTD_DIR/lib/libzstd.a" "$LIB_PATH" && \
    cp -f "$ZSTD_DIR/lib/zstd.h" "$INCLUDE_PATH" && \
    cp -f "$ZSTD_DIR/lib/zdict.h" "$INCLUDE_PATH" && \
    cp -f "$ZSTD_DIR/lib/zstd_errors.h" "$INCLUDE_PATH" || cleanAndExit

    echo "🚀 zstd installed"
}

function install_rocksdb(){
    echo "installing rocksdb"

    git clone --depth=1 https://github.com/facebook/rocksdb.git "$ROCKSDB_DIR" && cd "$ROCKSDB_DIR" || cleanAndExit

    make static_lib

    cp -fr "$ROCKSDB_DIR/include/rocksdb" "$INCLUDE_PATH" && \
    cp -f "$ROCKSDB_DIR/librocksdb.a" "$LIB_PATH" || cleanAndExit

    echo "🚀 rocksdb installed"
}

function main(){
    startup

    echo "🚧 start building library,this may takes a few time"

    instal_zlib

    intall_bzip2

    isntall_lz4

    install_snappy

    isntall_zstd

    install_rocksdb

    echo "🎉 build success! enjoy it!"

    cleanAndExit
}

main