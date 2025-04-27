// widget by Embed.im (Licenses & Credits: https://app.embed.im/licenses.txt)

// 尝试获取页面上已有的 sparkles 容器
var embedimSparkles = document.getElementById("embedim--sparkles");

// 如果容器不存在，则初始化整个效果
if (!embedimSparkles) {

    // 定义创建单个闪烁效果的函数
    let createSparkle = function(x, y) {
        // 创建一个新的 div 元素作为闪烁体
        const sparkle = document.createElement("div");

        // 从预定义颜色数组中随机选择一个颜色
        const color = embedimSparkleColors[Math.floor(Math.random() * embedimSparkleColors.length)];
        // 随机生成大小 (5px 到 20px 之间)
        const size = Math.random() * 15 + 5;
        // 随机生成动画的最终方向角度 (0 到 2*PI 弧度)
        const angle = Math.random() * Math.PI * 2;
        // 随机生成初始旋转角度 (-120deg 到 120deg)
        const rotation = Math.random() * 240 - 120;
        // 随机生成初始不透明度 (0.5 到 1.0)
        const opacity = Math.random() * 0.5 + 0.5;
        // 随机生成动画移动的距离 (10px 到 60px)
        const distance = Math.random() * 50 + 10;

        // --- 设置闪烁体的样式 ---
        sparkle.style.color = "#" + color; // 设置颜色 (SVG 使用 currentColor)
        sparkle.style.width = `${size}px`; // 设置宽度
        sparkle.style.height = `${size}px`; // 设置高度
        sparkle.style.left = `${x}px`;    // 设置初始水平位置 (鼠标 X)
        sparkle.style.top = `${y}px`;     // 设置初始垂直位置 (鼠标 Y)

        // 设置闪烁体的内部 HTML 为一个 SVG 图标
        // 注意：SVG 使用 currentColor 来继承 sparkle.style.color
        sparkle.innerHTML = `
            <svg xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="236.654 301.146 140.686 140.677">
                <path d="M302.184 304.627c1.547-4.642 8.101-4.642 9.648 0l9.072 27.244c4.037 12.138 13.573 21.66 25.725 25.71l27.23 9.072c4.641 1.547 4.641 8.102 0 9.649l-27.244 9.072a40.57 40.57 0 0 0-25.711 25.725l-9.072 27.229c-1.223 3.713-6.006 4.712-8.608 1.801a4.843 4.843 0 0 1-1.04-1.801l-9.072-27.243a40.611 40.611 0 0 0-25.711-25.711l-27.244-9.072c-3.713-1.224-4.726-6.006-1.814-8.608a4.889 4.889 0 0 1 1.814-1.041l27.244-9.072a40.675 40.675 0 0 0 25.711-25.71l9.072-27.244Z"/>
            </svg>
        `;

        // 将创建的闪烁体添加到容器中
        sparkleHolder.appendChild(sparkle);

        // --- 使用 Web Animations API 创建动画 ---
        sparkle.animate(
            [
                // 关键帧 1: 初始状态 (在鼠标位置，完全不透明，有初始旋转)
                {
                    opacity: opacity,
                    transform: `translate(-50%, -50%) scale(1) rotate(${rotation}deg)`
                },
                // 关键帧 2: 结束状态 (移动到随机位置，完全透明，无旋转，缩小到 0)
                {
                    opacity: 0,
                    transform: `translate(calc(-50% + ${Math.cos(angle) * distance}px), calc(-50% + ${Math.sin(angle) * distance}px)) scale(0) rotate(0)`
                }
            ],
            {
                duration: 1000, // 动画持续时间 1000ms (1秒)
                easing: "cubic-bezier(0.4, 0, 0.2, 1)", // 缓动函数
                fill: "forwards" // 动画结束后保持结束状态
            }
        );

        // 1 秒后移除闪烁体元素
        setTimeout(() => {
            sparkle.remove();
        }, 1000);
    };

    // 将内部函数赋值给外部变量 (这里看起来有些冗余，可以直接使用 createSparkle)
    // var embedimSparkle = createSparkle; // 保留原始逻辑，但 createSparkle 更清晰

    // --- 创建用于容纳所有闪烁效果的容器 ---
    const sparkleHolder = document.createElement("div");
    sparkleHolder.id = "embedim--sparkles"; // 给容器设置 ID，以便下次不再创建
    sparkleHolder.classList.add("embedim--sparkles"); // 添加 CSS 类名

    // 设置容器的内部 HTML，包含必要的 CSS 样式
    sparkleHolder.innerHTML = `
        <style>
            .embedim--sparkles {
                pointer-events: none; /* 容器不响应鼠标事件 */
                position: fixed;      /* 固定定位，覆盖整个视口 */
                top: 0;
                left: 0;
                bottom: 0;
                right: 0;
                z-index: 9999999999999; /* 非常高的 z-index，确保在顶层 */
            }
            .embedim--sparkles div { /* 容器内的每个闪烁体 (div) */
                position: absolute;   /* 绝对定位，相对于容器 */
                /* 初始尺寸，会被 JS 覆盖 */
                /* width: 10px; */
                /* height: 10px; */
            }
        </style>
    `;

    // 将容器添加到页面的 body 中
    document.body.appendChild(sparkleHolder);

    // --- 定义闪烁效果可用的颜色列表 ---
    const embedimSparkleColors = ["FFFFFF", "CCCCCC", "808080", "444444", "000000"];

    // --- 添加事件监听器 ---

    // 监听鼠标移动事件
    document.addEventListener("mousemove", e => {
        // 在鼠标当前位置创建闪烁效果
        createSparkle(e.clientX, e.clientY);
    });

    // 监听触摸移动事件 (针对移动设备)
    document.addEventListener("touchmove", e => {
        // 阻止默认的滚动行为 (如果需要的话)
        e.preventDefault();
        // 获取第一个触摸点
        const touch = e.touches[0];
        // 更新光标位置 (这里有个未定义的 'cursor' 变量，原始代码可能有误或依赖其他部分)
        // cursor.style.left = touch.clientX + "px";
        // cursor.style.top = touch.clientY + "px";
        // 在触摸点位置创建闪烁效果
        createSparkle(touch.clientX, touch.clientY);
    }, { passive: false }); // passive: false 允许 preventDefault 生效

} // 结束 if (!embedimSparkles) 块