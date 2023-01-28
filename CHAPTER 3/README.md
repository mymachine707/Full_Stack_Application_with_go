Working with Sessions, Error
    Handling, and Caching in Go
    In this chapter, we will cover the following recipes:
    Creating your first HTTP session
    {
        Bu faylda session libaryni yuklab oldim 
        bu code login qilinganda cookie faylga vaqtinchalik id saqlab turadi page to'liq dostup beradi
        login qilingandan keyin home page ochiladi
        logout qilingan paytda dostup yopiladi

    }
    Managing your HTTP session using Redis
    {
        ishlamadi
    }
    Creating your first HTTP cookie 
    {
        
    }
    Implementing caching in Go
    Implementing HTTP error handling in Go
    Implementing login and logout in a web application

Ba'zan biz seanslar va cookie-fayllar yordamida osongina erishish mumkin bo'lgan ma'lumotlar bazasida saqlash o'rniga, foydalanuvchi ma'lumotlari kabi ma'lumotlarni ilova darajasida saqlashni xohlaymiz. Ularning orasidagi farq shundaki, seanslar server tomonida, kukilar esa mijoz tomonida saqlanadi. Shuningdek, maʼlumotlar bazasiga yoki veb-xizmatga keraksiz qoʻngʻiroqlarni oldini olish uchun statik maʼlumotlarni keshlashimiz va veb-ilovani ishlab chiqishda xatolarni qayta ishlashni amalga oshirishimiz kerak boʻlishi mumkin. Ushbu bobda ko'rib chiqilgan tushunchalarni bilgan holda, biz ushbu funktsiyalarning barchasini juda oson tarzda amalga oshirishimiz mumkin. Ushbu bobda biz HTTP seansini yaratishdan boshlaymiz, keyin Redis yordamida uni qanday boshqarishimiz, cookie-fayllar yaratish, HTTP javoblarini keshlash, xatolarni qayta ishlashni amalga oshirish va oxir-oqibat Go-da kirish va chiqish mexanizmlarini amalga oshirish bilan yakunlaymiz.