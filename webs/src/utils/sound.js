const allSound = [
    "wow.mp3",
    "manbo.wav",
    "aminuosi.m4a",
    "oye.mp3",
    "siren.mp3",
    "duang.m4a",
]

export function playSound() {
    const file = allSound[Math.floor(Math.random() * allSound.length)]
    const audio = new Audio(`/sounds/${file}`)
    audio.play().catch((e) => {
        console.log("播放声音失败", e)
    })
}