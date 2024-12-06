/*顶部导航栏跳转到对应位置*/
function scrollToSection(sectionId) {
    const targetSection = document.getElementById(sectionId);
    window.scrollTo({
        top: targetSection.offsetTop,
        behavior: 'smooth'  // 设置平滑滚动
    });
}


/*跳出弹窗*/
// 确保脚本在 DOM 加载完成后执行
document.addEventListener("DOMContentLoaded", () => {
    const emailButton = document.getElementById("email-button");
    const emailPopup = document.getElementById("email-popup");
    const closePopupButton = document.getElementById("close-popup");

    // 显示弹窗
    emailButton.addEventListener("click", (event) => {
        event.preventDefault(); // 阻止默认跳转行为
        emailPopup.classList.remove("hidden"); // 显示弹窗
        navigator.clipboard.writeText("newthread_geek@outlook.com").then(() => {
            console.log("邮箱地址已复制到剪贴板");
        });
    });

    // 关闭弹窗
    closePopupButton.addEventListener("click", () => {
        emailPopup.classList.add("hidden"); // 隐藏弹窗
    });
});


/*问题回答*/
// 页面加载后执行
document.addEventListener("DOMContentLoaded", () => {
    // 获取所有问题和回答元素
    const questions = document.querySelectorAll(".QA-question");
    const answers = document.querySelectorAll(".QA-answer");

    // 初始隐藏所有回答
    answers.forEach((answer) => {
        answer.style.display = "none"; // 默认隐藏
    });

    // 为每个问题绑定点击事件
    questions.forEach((question, index) => {
        question.addEventListener("click", () => {
            // 当前问题的回答
            const answer = answers[index];

            // 收起其他回答
            answers.forEach((otherAnswer, i) => {
                if (i !== index) {
                    otherAnswer.style.display = "none"; // 隐藏其他回答
                }
            });

            // 切换当前回答的显示状态
            answer.style.display = (answer.style.display === "block") ? "none" : "block";
        });
    });
});


/*火箭返回按钮*/
document.addEventListener("DOMContentLoaded", () => {
    const pointer = document.getElementById("pointer");

    if (pointer) {
        // 滚动事件
        window.addEventListener("scroll", () => {
            const scrollPosition = window.scrollY;
            pointer.style.display = scrollPosition >= 800 ? "block" : "none";
        });

        // 点击事件
        pointer.addEventListener("click", () => {
            window.scrollTo({ top: 0, behavior: "smooth" });
        });
    }
});
