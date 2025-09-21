const allSound = [
    "wow.mp3",
    "manbo.wav",
    "aminuosi.m4a",
]

export function playSound() {
    const file = allSound[Math.floor(Math.random() * allSound.length)]
    const audio = new Audio(`/sounds/${file}`)
    audio.play().catch((e) => {
        console.log("播放声音失败", e)
    })
}