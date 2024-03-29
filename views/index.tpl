<!DOCTYPE html>
<html>
        <head>
                <title> 短链接_稳定、强大、免费短网址服务</title>
                <!-- 新 Bootstrap 核心 CSS 文件 -->
                        <link rel="stylesheet" href="//cdn.bootcss.com/bootstrap/3.3.5/css/bootstrap.min.css">
                        <!-- 可选的Bootstrap主题文件（一般不用引入） -->
                        <link rel="stylesheet" href="//cdn.bootcss.com/bootstrap/3.3.5/css/bootstrap-theme.min.css">
                        <!-- jQuery文件。务必在bootstrap.min.js 之前引入 -->
                        <script src="//cdn.bootcss.com/jquery/1.11.3/jquery.min.js"></script>
                        <!-- 最新的 Bootstrap 核心 JavaScript 文件 -->
                        <script src="//cdn.bootcss.com/bootstrap/3.3.5/js/bootstrap.min.js"></script>

                        <link rel="stylesheet" type="text/css" href="/static/css/style.css">

                        
        </head>
        <body>

          <div class="main">        
                  <h4 id="title">短链接服务</h4>
            <form>
                      <div class="input-group">
                          <input type="text" class="form-control" type="text" name="source" placeholder="url input here">
                          <span class="input-group-btn">
                                    <button class="btn btn-primary" id="click" type="button">生成短链接</button>
                          </span>
                        </div>
            </form>
                  <div class="show">
                   <div>
                                 短链为:  <a id="result" ></a>
                   </div>           
                <div>
                   <a id="message" ></a>
                </div>
                 </div>
        </div>

        <script>
                $('#click').click(function(){
                        empty();
                        var sourceUrl = $('input[name=source]').val();
                        // console.log(sourceUrl);
                        $.get('/api/url', {'sourceUrl':sourceUrl},function(data){
                                if(data.success){
                                        var shortUrl = window.location.host + "/" + data.result.ShortUrl;
                                        $('#result').html(shortUrl);
                                        $('#result').attr("href","http://"+shortUrl);
                                }else{
                                        $('#message').html(data.message);
                                }
                        });
                });

                function empty(){
                        $('#message').html("");
                        $('#result').html("");
                }
        </script>
        </body>
</html>