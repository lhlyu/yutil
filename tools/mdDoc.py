# encoding=utf-8

# 根据 go 文件生产 markdown 文件
# 目前只针对 func 函数
# 过滤xxx_test.go的文件

from jinja2 import Environment, FileSystemLoader
import os

env = Environment(loader = FileSystemLoader("./"))
template = env.get_template('tmp.md')

def create(fileName,content):
    dir = os.path.abspath(os.path.dirname(os.getcwd()))
    filePath = os.path.join(dir,"doc",fileName)
    with open(filePath,"w",encoding="utf-8") as f:
        f.write(content)
    pass

def get_all_file():
    dir = os.path.abspath(os.path.dirname(os.getcwd()))
    fileList = []
    for root, dirs, files in os.walk(dir):
        # print(root) #当前目录路径
        # print(dirs) #当前路径下所有子目录
        # print(files) #当前路径下所有非目录子文件
        for f in files:
            if f[-3:] == ".go" and f[-8:] != "_test.go":
                fileList.append(os.path.join(root,f))
    if len(fileList) == 0:
        return
    get_func_arr(fileList)

def get_func_arr(fileList):
    for v in fileList:

        Titile = ""
        funcArr = []
        name = os.path.basename(v)[:-3]
        Titile = name.title()
        with open(v,"r",encoding="utf-8") as f:
            lines = f.readlines()
            for index,line in enumerate(lines):
                line = line.strip()
                if line.find("// @") == 0:
                    funcDict = {}
                    funcDict["desc"] = line[4:]
                    next = lines[index + 1].strip()
                    funcDict["func"] = next[5:-1]
                    funcArr.append(funcDict)
                    package = lines[0].strip()
                    get_content(Titile,package,funcArr)

    pass

def get_content(title,package,funcArr):
    content = template.render(Title=title,package=package, funcArr=funcArr)
    create(title.lower() + ".md",content)
    pass

def run():
    get_all_file()
    pass

if __name__ == '__main__':
    run()