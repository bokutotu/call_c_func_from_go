cmake_minimum_required(VERSION 3.9)
set(LIBRARY_NAME hoge)

set(LIBRARY_OUTPUT_PATH ${CMAKE_SOURCE_DIR}/bin)

project(${LIBRARY_NAME} VERSION 1.0.1 DESCRIPTION "mylib description")

include(GNUInstallDirs)
add_library(${LIBRARY_NAME} SHARED src/hoge.c)
set_target_properties(${LIBRARY_NAME} PROPERTIES
    VERSION ${PROJECT_VERSION}
    SOVERSION 1
    PUBLIC_HEADER include/hoge.h)
target_include_directories(${LIBRARY_NAME} PRIVATE .)
