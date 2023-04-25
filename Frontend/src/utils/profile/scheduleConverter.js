// Convert the schedules from the backend to the frontend format
export const convertBackendSchedulesToFrontend = (schedules, isMerge=false) => {
  function capitalizeFirstLetter(string) {
    return string.charAt(0).toUpperCase() + string.slice(1);
  }
  const days = [
    "sunday",
    "monday",
    "tuesday",
    "wednesday",
    "thursday",
    "friday",
    "saturday",
  ];
  // schedules is an object with key as day and value as an array of 24 hours as boolean
  const convertedSchedules = days.map((day) => {
    let timeSlots = []
    for (let i = 0; i < 24; i++) {
      let from;
      let to;
      if (schedules[day][i]) {
        from = (i < 10 ? "0" : "") + i + ":00"
        to = (i + 1 < 10 ? "0" : "") + (i + 1) + ":00"
        // merge if isMerge is true
        while(isMerge && i < 24 && schedules[day][i]) {
          to = (i + 1 < 10 ? "0" : "") + (i + 1) + ":00"
          i++;
        }
        timeSlots.push({from, to})
      }
    }
    const filteredTimeSlots = timeSlots.filter(
      (timeSlot) => timeSlot !== undefined
    );
    return {
      day: capitalizeFirstLetter(day),
      timeSlot: (filteredTimeSlots || []),
    };
  });
  return convertedSchedules;
};
// Convert the schedules from the frontend to the backend format
export const convertFrontendSchedulesToBackend = (schedules) => {
  const dayMapper = {
    Sunday: 0,
    Monday: 1,
    Tuesday: 2,
    Wednesday: 3,
    Thursday: 4,
    Friday: 5,
    Saturday: 6,
  };
  // {from:"04:00", to:"06:00"} => [4,5]
  const timeSlotMapper = ({ from, to }) => {
    const fromHour = parseInt(from.split(":")[0]);
    const toHour = parseInt(to.split(":")[0]);
    const timeSlot = [];
    for (let i = fromHour; i < toHour; i++) {
      timeSlot.push(i);
    }
    return timeSlot;
  };
  // {day: "Monday", timeSlot: [{from:"04:00", to:"06:00"}]} => {day: 1, hour:4}, {day: 1, hour:5}
  const convertedSchedules = schedules
    .map((schedule) => {
      const hours = schedule.timeSlot.map(timeSlotMapper);
      const day = dayMapper[schedule.day];
      const convertedHours = hours.flat().map((hour) => ({ day, hour }));
      return convertedHours;
    })
    .flat();
  return convertedSchedules;
};
