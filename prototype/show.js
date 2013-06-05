var show = (function(document,undefined){
    var self ={};
    
    self.slides = {};
    self.index = 0;
    
    self.init = function(){
        var slides = this.slides.slides;
        var title, content, section;
        var container = document.getElementById("slides");
        for(var i = 0; i < slides.length; i++){
                s = slides[i];
                title = document.createElement("h2");
                content = document.createElement("div");
                section = document.createElement("section");
                    
                title.innerHTML = s.title;
                content.innerHTML = s.content;
                    
                section.appendChild(title);
                section.appendChild(content);
                container.appendChild(section);
                
                section.style.webkitAnimationPlayState = "paused";
            
                if(i == 0){
                    section.classList.add("present");
                }else{
                    section.classList.add("future");
                }
            
                s.e = section;
        }
    };
    
    var jump = function(index){
        var last = self.index;
        
        if(index == last){
            return;
        }
        
        if(index > self.slides.slides.length){
            return
        }
        
        var now = self.slides.slides[self.index].e;
        
        self.index = index;
        var next = self.slides.slides[index].e;
        
        for(var i = self.index; i < self.slides.slides.length; i++){
            self.slides.slides[i].e.classList.remove("future");
            self.slides.slides[i].e.classList.remove("present");
            self.slides.slides[i].e.classList.remove("past");
            self.slides.slides[i].e.classList.add("future");
        }
        
        for(var i = 0; i < index; i++){
            console.log(self.slides.slides[i])
            self.slides.slides[i].e.classList.remove("future");
            self.slides.slides[i].e.classList.remove("present");
            self.slides.slides[i].e.classList.remove("past");
            self.slides.slides[i].e.classList.add("past");
        }
        
        next.classList.remove("future");
        next.classList.remove("past");
        next.classList.remove("present");
        
        now.classList.remove("future");
        now.classList.remove("past");
        now.classList.remove("present");
        
        
        now.classList.add("present");
        
        
        
        if(last < index){
            next.classList.add("future")
            now.addEventListener("webkitAnimationEnd", function(){
            this.style.webkitAnimationPlayState = "paused";
            this.classList.add("past");
            this.classList.remove("present");
            this.removeEventListener("webkitAnimationEnd");
        });
        
        next.addEventListener("webkitAnimationEnd", function(){
            this.style.webkitAnimationPlayState = "paused";
            this.classList.add("present");
            this.classList.remove("future");
            this.removeEventListener("webkitAnimationEnd");
        });
        
        //now.classList.add("leftFadeOut");
        //next.classList.add("leftFadeIn");
        
        
        now.style.webkitAnimationPlayState = "running";
        next.style.webkitAnimationPlayState = "running";
        
        now.addEventListener("webkitAnimationStart", function(){
            this.style.opacity = 0;
        });
        
        next.addEventListener("webkitAnimationStart", function(){
            this.style.opacity = 1;
        });
                
                return;
                
        }
        
        next.classList.add("past")
            now.addEventListener("webkitAnimationEnd", function(){
                this.classList.add("future");
                this.style.webkitAnimationPlayState = "paused";    
                this.classList.remove("present");
                    this.removeEventListener("webkitAnimationEnd");
                });
        
            next.addEventListener("webkitAnimationEnd", function(){
                this.classList.add("present");
                this.style.webkitAnimationPlayState = "paused";
                this.classList.remove("past");
                this.removeEventListener("webkitAnimationEnd");
            });
        
            now.style.webkitAnimationPlayState = "running";
            next.style.webkitAnimationPlayState = "running";
        
            now.addEventListener("webkitAnimationStart", function(){
                this.style.opacity = 0;
            });
        
            next.addEventListener("webkitAnimationStart", function(){
                this.style.opacity = 1;
            });
    };
    
    var input = document.getElementById("panel");
    input.childNodes[1].addEventListener("keydown", function(event){
        if(event.keyCode === 13){
            var v = this.value;
            this.value = "";
            input.style.opacity = 0;
            this.blur();
            jump(parseInt(v));
        }
    });
    
    
    document.body.addEventListener("keydown", function(event){
        var c = event.keyCode;
        if(c === 37 || c === 39 || c === 32)event.preventDefault();
    });
    
    document.body.addEventListener("keyup", function(event){
        var c = event.keyCode;
        
        /*left*/
        if(c === 37){
            jump(self.index-1);
            event.preventDefault();
            return
        }
        
        /*Right or Spacebar*/
        if(c === 39 || c === 32){
            jump(self.index+1);
            event.preventDefault();
            return
        }
        
        if(c === 190 && event.shiftKey){
            var l = input.style.opacity^1;
            
            input.style.opacity = l;
            if(l== 1){
                //input.input.focus();
                input.childNodes[1].focus();
            }
        }
    });
    
    
    
    return self;
}(document));
