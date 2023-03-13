// Convert the schedules from the backend to the frontend format
export const convertBackendSchedulesToFrontend = (schedules) => {
    const dayMapper = {
        0: "Sunday",
        1: "Monday",
        2: "Tuesday",
        3: "Wednesday",
        4: "Thursday",
        5: "Friday",
        6: "Saturday",
    };
    const timeSlotMapper = {}
    for (let i = 0; i < 24; i++) {
        if (i < 10) {
            if (i === 9) timeSlotMapper[i] = {from: `0${i}:00`, to: `${i+1}:00`}
            else {
                timeSlotMapper[i] = {from: `0${i}:00`, to: `0${i+1}:00`}
            }
        }else{
            timeSlotMapper[i] = {from: `${i}:00`, to: `${(i+1)%24}:00`}
        }
    }
    const convertedSchedules = schedules.map((schedule) => {
        const timeSlot = schedule.hour.map((time) => timeSlotMapper[time])
        return {day: dayMapper[schedule.day], timeSlot}
    })
    return convertedSchedules
}
// Convert the schedules from the frontend to the backend format
export const convertFrontendSchedulesToBackend = (schedules) => {
    const dayMapper = {
        "Sunday": 0,
        "Monday": 1,
        "Tuesday": 2,
        "Wednesday": 3,
        "Thursday": 4,
        "Friday": 5,
        "Saturday": 6,
    };
    // {from:"04:00", to:"06:00"} => [4,5]
    const timeSlotMapper = ({from, to}) => {
        const fromHour = parseInt(from.split(":")[0])
        const toHour = parseInt(to.split(":")[0])
        const timeSlot = []
        for (let i = fromHour; i < toHour; i++) {
            timeSlot.push(i)
        }
        return timeSlot
    }
    // {day: "Monday", timeSlot: [{from:"04:00", to:"06:00"}]} => {day: 1, hour:4}, {day: 1, hour:5}
    const convertedSchedules = schedules.map((schedule) => {
        const hours = schedule.timeSlot.map(timeSlotMapper)
        const day = dayMapper[schedule.day]
        const convertedHours = hours.flat().map(hour => ({day, hour}))
        return convertedHours
    }).flat()
    return convertedSchedules
}