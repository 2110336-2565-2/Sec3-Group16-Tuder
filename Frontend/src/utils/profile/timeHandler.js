import { toast } from "react-hot-toast";

export const adjustFormat = (time) => {
    // mm must be 00
    if (time && time.split(":")[1] !== "00") {
      toast.error("Time must be in 00 minutes");
      // set mm to 00
      time = time.split(":")[0] + ":00";
    }
    return time;
};

export const isOverlapped = (timeSlotToAdd, timeSlot) => {
  if (timeSlotToAdd.from === timeSlotToAdd.to) {
    toast.error("Minimum time slot is 1 hour");
    return true;
  }
  // Check if time slot is already added by check intersection
  const overlappedSlots = timeSlot.filter(
    (time) => (timeSlotToAdd.from >= time.from && timeSlotToAdd.from < time.to) || 
    (timeSlotToAdd.to > time.from && timeSlotToAdd.to <= time.to) ||
    (timeSlotToAdd.from <= time.from && timeSlotToAdd.to >= time.to)
  );
  if (overlappedSlots.length > 0) {
    toast.error("Time slot is already added");
    return true;
  }
  return false;
};

export const getMergedTimeSlot = (timeSlot, timeSlotToAdd) => {
  const mergedTimeSlot = [...timeSlot, timeSlotToAdd];
    mergedTimeSlot.sort((a, b) => {
      const aMinutes = parseInt(a.from.split(':')[0], 10) * 60 + parseInt(a.from.split(':')[1], 10);
      const bMinutes = parseInt(b.from.split(':')[0], 10) * 60 + parseInt(b.from.split(':')[1], 10);
      return aMinutes - bMinutes;
    });
    let j = mergedTimeSlot.length - 1;
    while (j > 0) {
      let i = j - 1;
      console.log("i",mergedTimeSlot[i],"j",mergedTimeSlot);
      if (mergedTimeSlot[j].from == mergedTimeSlot[i].to) {
        mergedTimeSlot[i].to = mergedTimeSlot[j].to;
        // Remove the overlapped time slot
        mergedTimeSlot.splice(j, 1);
      }
      j--;
    }
    return mergedTimeSlot;
}