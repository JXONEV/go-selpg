# go-selpg

    1.该命令将把“in.txt”的第 1 页写至标准输出。
![](https://github.com/JXONEV/go-selpg/raw/master/image/2.png)



    2.selpg 将第1页写至标准输出；标准输出被 shell／内核重定向至“output”。
![](https://github.com/JXONEV/go-selpg/raw/master/image/3.png)

![](https://github.com/JXONEV/go-selpg/raw/master/image/4.png)

    3.selpg 将第 10 页到第 20 页写至标准输出（屏幕）；所有的错误消息被 shell／内核重定向至“error_file”。请注意：在“2”和“>”之间不能有空格；这是 shell 语法的一部分（请参阅“man bash”或“man sh”）。
![](https://github.com/JXONEV/go-selpg/raw/master/image/8.png)

![](https://github.com/JXONEV/go-selpg/raw/master/image/9.png)

    4.selpg 将第 10 页到第 20 页写至标准输出，标准输出被重定向至“output_file”；selpg 写至标准错误的所有内容都被重定向至“error_file”。当“input_file”很大时可使用这种调用；您不会想坐在那里等着 selpg 完成工作，并且您希望对输出和错误都进行保存。
![](https://github.com/JXONEV/go-selpg/raw/master/image/5.png)
    此时error_file中为空（没有错误信息

![](https://github.com/JXONEV/go-selpg/raw/master/image/6.png)
    此时error_file中有错误信息
![](https://github.com/JXONEV/go-selpg/raw/master/image/7.png)

