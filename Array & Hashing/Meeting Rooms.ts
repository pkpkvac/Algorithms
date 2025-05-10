// Given an array of meeting time interval objects consisting of start and end times
//  [[start_1,end_1],[start_2,end_2],...] (start_i < end_i),
// determine if a person could add all meetings to their schedule without any conflicts.

// Example 1:
// Input: intervals = [(0,30),(5,10),(15,20)]
// Output: false
// Explanation:
// (0,30) and (5,10) will conflict
// (0,30) and (15,20) will conflict

// Example 2:
// Input: intervals = [(5,8),(9,15)]
// Output: true

// Note:
// (0,8),(8,10) is not considered a conflict at 8

class Interval {
	start: number;
	end: number;

	constructor(start: number, end: number) {
		this.start = start;
		this.end = end;
	}
}

function canAttendMeetings(meetings: Interval[]): boolean {
	// sort meetings by their start

	meetings.sort((a, b) => a.start - b.start);

	for (let i = 0; i < meetings.length - 1; i++) {
		if (meetings[i].end > meetings[i + 1].start) {
			return false;
		}
	}

	console.log(meetings);
	return true;
}

let i1 = new Interval(5, 8);
let i2 = new Interval(9, 15);

let meetingList = [i1, i2];

console.log(canAttendMeetings(meetingList));
