
# go-selpg

环境：Win10 

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

    5.selpg 将第 10 页到第 20 页写至标准输出，标准输出被重定向至“output_file”；selpg 写至标准错误的所有内容都被重定向至 /dev/null（空设备），这意味着错误消息被丢弃了。设备文件 /dev/null 废弃所有写至它的输出，当从该设备文件读取时，会立即返回 EOF。
      
      
    6.该命令将页长设置为 2 行，这样 selpg 就可以把输入当作被定界为该长度的页那样处理。第 1 页到第 3 页被写至 selpg 的标准输出（屏幕）。
![](https://github.com/JXONEV/go-selpg/raw/master/image/10.png)

    7.假定页由换页符定界。第 1 页到第 3 页被写至 selpg 的标准输出（屏幕）。
![](https://github.com/JXONEV/go-selpg/raw/master/image/11.png)

    8.该命令利用了 Linux 的一个强大特性，即：在“后台”运行进程的能力。在这个例子中发生的情况是：“进程标识”（pid）如 1234 将被显示，然后 shell 提示符几乎立刻会出现，使得您能向 shell 输入更多命令。同时，selpg 进程在后台运行，并且标准输出和标准错误都被重定向至文件。这样做的好处是您可以在 selpg 运行时继续做其它工作。
![](https://github.com/JXONEV/go-selpg/raw/master/image/12.png)
